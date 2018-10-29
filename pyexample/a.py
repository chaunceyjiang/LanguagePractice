class Person(object):
    def __init__(self, name, gender):
        self.name = name
        self.gender = gender
    def whoAmI(self):
        return 'I am a Person, my name is %s' % self.name

class Student(Person):
    def __init__(self, name, gender, score):
        super(Student, self).__init__(name, gender)
        self.score = score
    def whoAmI(self):
        return 'I am a Student, my name is %s' % self.name

class Teacher(Person):
    def __init__(self, name, gender, course):
        super(Teacher, self).__init__(name, gender)
        self.course = course
    def whoAmI(self):
        return 'I am a Teacher, my name is %s' % self.name

def who_am_i(x):
    print x.whoAmI

p = Person('Tim', 'Male')
s = Student('Bob', 'Male', 88)
t = Teacher('Alice', 'Female', 'English')
# Person.whoAmI()
p.whoAmI

who_am_i(Person)
who_am_i(s)
who_am_i(t)
# who_am_i(1)
import psutil
print psutil.net_if_addrs()
from itertools import ifilter
a={"a":1,"b":2,"c":3}
# a=["a","b","c"]

def f(x):
    print x=="a"
    return x!="a"

a=dict(filter(lambda x:"a" not in x,a.items()))
print a