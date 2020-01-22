package main

import (
	"strings"
	"sync"
)

// 基于消息队列与2PC配合延迟队列完成消息事务
// http://www.sreguide.com/%E5%88%86%E5%B8%83%E5%BC%8F/22.html

const (
	// TaskPrefix 任务key的前缀
	TaskPrefix string = "task-"
	// CommitTaskPrefix 提交任务的前缀
	CommitTaskPrefix string = "commit-"
	// ClearTaskPrefix 清除任务
	ClearTaskPrefix string = "clear-"
)

type Task struct {
	Name   string            // 任务名字
	Group  string            // 任务所在的任务组
	Config map[string]string //任务的配置
}

func (t *Task) updateConfig(config map[string]string) map[string]string {
	// TODO
	return config
}

type Event struct {
	Key   string
	Name  string
	Value interface{}
}

type EventListener interface {
	onEvent(event *Event)
}

type MemoryQueue struct {
	done     chan struct{}
	queue    chan Event
	listener []EventListener
	wg       sync.WaitGroup
}

func (mq *MemoryQueue) Push(eventType, name string, value interface{}) {
	mq.queue <- Event{
		Key:   eventType + name,
		Name:  name,
		Value: value,
	}
	mq.wg.Add(1)
}

func (mq *MemoryQueue) AddListener(listener EventListener) bool {
	for _, l := range mq.listener {
		if l == listener {
			return false
		}
	}
	mq.listener = append(mq.listener, listener)
	return true
}

func (mq *MemoryQueue) Notify(event *Event) {
	for _, l := range mq.listener {
		l.onEvent(event)
	}
}

func (mq *MemoryQueue) poll() {
	for {
		select {
		case <-mq.done:
			break
		case e := <-mq.queue:
			mq.Notify(&e)
		}
	}
}
func (mq *MemoryQueue) Start() {
	go mq.poll()
}
func (mq *MemoryQueue) Stop() {
	mq.wg.Wait()
	close(mq.queue)
}
func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{
		done:     make(chan struct{}),
		queue:    make(chan Event),
		listener: make([]EventListener, 0),
		wg:       sync.WaitGroup{},
	}
}

type ConfigUpdateCallback func()
type Worker struct {
	name string
	// 本地延迟存储
	deferredTaskUpdates map[string][]Task
	onCommit            ConfigUpdateCallback
}

func (w *Worker) onTaskClear(event *Event) {
	task, ok := event.Value.(Task)
	if !ok {
		return
	}
	_, found := w.deferredTaskUpdates[task.Group]
	if !found {
		return
	}
	delete(w.deferredTaskUpdates, task.Group)
}

func (w *Worker) onTaskCommit(event *Event) {
	tasks, found := w.deferredTaskUpdates[event.Name]
	if !found {
		return
	}

}

func (w *Worker) getTaskConfig(tasks []Task) map[string]string {
	config := make(map[string]string)
	for _, t := range tasks {
		// 刷新配置
		config = t.updateConfig(config)
	}
	return config
}

func (w *Worker) onEvent(event *Event) {
	switch {
	case strings.Contains(event.Key, TaskPrefix):

	}
}

func main() {

}
