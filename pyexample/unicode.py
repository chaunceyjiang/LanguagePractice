class A:
    def __init__(self, a):
        self.a=a
    def __unicode__(self):
         return 'hello' + self.a

a=A('4')
print("2"+str(a)) 
