package main

import "fmt"

// CircularIterator 环状迭代器
type CircularIterator struct {
	list []interface{}
	next int
}

// Next 返回下一个元素
func (c *CircularIterator) Next() interface{} {
	item := c.list[c.next]
	c.next = (c.next + 1) % len(c.list)
	return item
}

// Add 添加任务
func (c *CircularIterator) Add(v interface{}) bool {
	for _, item := range c.list {
		if v == item {
			return false
		}
	}
	c.list = append(c.list, v)
	return true
}

// Member 任务组成员
type Member struct {
	id    int
	tasks []*Task
}

// ID 返回当前memberID
func (m *Member) ID() int {
	return m.id
}

// AddTask 为member添加任务
func (m *Member) AddTask(t *Task) bool {
	for _, task := range m.tasks {
		if task == t {
			return false
		}
	}
	m.tasks = append(m.tasks, t)
	return true
}

// Execute 执行任务
func (m *Member) Execute() {
	for _, task := range m.tasks {
		fmt.Printf("Member %d run task %s\n", m.ID(), task.Execute())
	}
}

// Task 任务
type Task struct {
	name string
}

// Execute 执行task返回结果
func (t *Task) Execute() string {
	return "Task " + t.name + " run success"
}

// Coordinator 协调者
type Coordinator struct {
	members []*Member
	tasks   []*Task
}

// TaskAssignments 为member分配任务
func (c *Coordinator) TaskAssignments() map[int]*Member {
	taskAssignments := make(map[int]*Member)

	// 构建迭代器
	memberIt := c.getMemberIterator()
	for _, task := range c.tasks {
		member := memberIt.Next().(*Member)

		_, err := taskAssignments[member.ID()]
		if err == false {
			taskAssignments[member.ID()] = member
		}
		member.AddTask(task)
	}

	return taskAssignments
}

func (c *Coordinator) getMemberIterator() *CircularIterator {
	// 通过当前成员, 构造成员队列
	members := make([]interface{}, len(c.members))
	for index, member := range c.members {
		members[index] = member
	}

	return NewCircularIterftor(members)
}

// AddMember 添加member组成员
func (c *Coordinator) AddMember(m *Member) bool {
	for _, member := range c.members {
		if member == m {
			return false
		}
	}
	c.members = append(c.members, m)
	return true
}

// AddTask 添加任务
func (c *Coordinator) AddTask(t *Task) bool {
	for _, task := range c.tasks {
		if task == t {
			return false
		}
	}
	c.tasks = append(c.tasks, t)
	return true
}

// NewCircularIterftor 返回迭代器
func NewCircularIterftor(list []interface{}) *CircularIterator {
	iterator := CircularIterator{}
	for _, item := range list {
		iterator.Add(item)
	}
	return &iterator
}

// NewCoordinator 返回协调器
func NewCoordinator() *Coordinator {
	c := Coordinator{}
	return &c
}

func main() {
	coordinator := NewCoordinator()
	for i := 0; i < 10; i++ {
		m := &Member{id: i}
		coordinator.AddMember(m)
	}

	for i := 0; i < 30; i++ {
		t := &Task{name: fmt.Sprintf("task %d", i)}
		coordinator.AddTask(t)
	}

	result := coordinator.TaskAssignments()
	for _, member := range result {
		member.Execute()
	}
}