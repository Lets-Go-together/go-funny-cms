package schedule

import (
	"encoding/json"
	"github.com/blinkbean/dingtalk"
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"time"
)

var (
	SCHEDULE_KEY = "SCHEDULE:JOBS"

	STATUS_RUNING   = 1
	STATUS_STARTING = 2
	STATUS_STOPPING = 3
	STATUS_STOPPED  = 4
)

type Process struct {
	Name    string       `json:"name"`
	Content string       `json:"content"`
	Spec    string       `json:"spec"`
	TimeAt  string       `json:"time_at"`
	EntryId cron.EntryID `json:"entry_id"`
	Status  int          `json:"status"`
	StopAt  string       `json:"stop_at"`
}

type Schedule struct {
	cron    *cron.Cron
	client  *redis.Client
	Content string `json:"content"`
}

// InitSchedule 初始化队列管理器
func InitSchedule() {
	var schedule Schedule
	schedule.cron = cron.New()
	schedule.client = config.Redis

	// 运行一下假数据
	DispatchTestProcess()

	schedule.RunJobs()
	schedule.cron.AddFunc("* * * * *", schedule.ManangeJob)
	schedule.cron.Run()
}

func DispatchTestProcess() {
	process := Process{
		Name:    "每一个小时运行一次",
		Content: "每一个小时运行一次",
		Spec:    "*/60 * * * *",
		Status:  STATUS_STARTING,
	}
	Dispatch(process)
}

// ManangeJob 维护和管理每个任务的执行停止删除
func (that Schedule) ManangeJob() {
	that.RunJobs()
}

// GetJobs 从 redis 获取当前已有的任务
func (that Schedule) GetJobs() []Process {
	var jobs []Process
	if r, err := that.client.HGetAll(SCHEDULE_KEY).Result(); err == nil {
		for _, item := range r {
			var process Process
			json.Unmarshal([]byte(item), &process)
			jobs = append(jobs, process)
		}
	}

	return jobs
}

// RunJobs 从 redis 执行当前已有的任务
func (that Schedule) RunJobs() {
	jobs := that.GetJobs()
	timeAt := time.Now().Format("2006-01-02 15:04:05")
	for _, item := range jobs {
		switch item.Status {
		case STATUS_STOPPING:
			that.cron.Remove(item.EntryId)
			item.Status = STATUS_STOPPED
			item.StopAt = timeAt
		case STATUS_RUNING:
			// 如果检测到当前的entry ID 不存在时，我们对他进行重启
			if entry := that.cron.Entry(item.EntryId); !entry.Valid() {
				item.EntryId = that.StartJob(item)
				item.Status = STATUS_RUNING
				item.TimeAt = timeAt
				logger.Info(item.Name, "重启中... | Time: "+timeAt)
				continue
			}
			logger.Info(item.Name, "正常运行中... | Time: "+timeAt)
		case STATUS_STARTING:
			entry_id := that.StartJob(item)
			logger.Info(item.Name, "启动完成... | Time: "+timeAt)
			item.Status = STATUS_RUNING
			item.EntryId = entry_id
			item.TimeAt = timeAt
		case STATUS_STOPPED:
			logger.Info(item.Name, "检测到已停止... | Time: "+timeAt)
		default:
			logger.Error(item.Name, "检测异常... | Time: "+timeAt)
		}

		// 同步到redis中
		Dispatch(item)
	}
}

// 添加一个运行队列
func (that Schedule) StartJob(process Process) cron.EntryID {
	that.Content = process.Content
	entry_id, _ := that.cron.AddFunc(process.Spec, that.DingTalk)

	return entry_id
}

// 用来执行队列做的事情
func (that Schedule) DingTalk() {
	var dingToken = []string{"b960b8c2240b7d0f05b1ffbf26b4a7807efa2fb22603127dcbdc618ea48607ea"}
	cli := dingtalk.InitDingTalk(dingToken, "任务")
	cli.SendTextMessage(that.Content + " | TimeAt:" + time.Now().Format("2006-01-02 15:04:05"))
}

// 分发任务
func Dispatch(process Process) {
	r, _ := json.Marshal(process)
	if _, err := config.Redis.HSet(SCHEDULE_KEY, process.Name, string(r)).Result(); err != nil {
		logger.Error("REDIS ERROR", err.Error())
	}
}
