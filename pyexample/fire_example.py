from __future__ import print_function
import fire


class Main():

    def printf(self):
        print("yes")
    def help(self):
        print("USAGE ")

    def register_dp(self,a,b,c):
        """
        打印信息
        """
        print("register_dp",a,b,c)
if __name__=="__main__":
    fire.Fire(Main)


    