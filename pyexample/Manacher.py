# coding: utf-8
'''
最长回文子串
'''


s= "ccc"
class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        return self.manacher(s)
        
    def pre_process(self,seq):
        res = ['#{}'.format(elem) for elem in seq]
        res.append('#$')
        res.insert(0,'^')
        return ''.join(res)
    def bef_process(self,max_m,index,T):
        S = []
        for i in range(index-max_m,index+max_m):
            S.append(T[i])
        while '#' in S:
            S.remove('#')
        return ''.join(S)

    def manacher(self,seq):
        T = self.pre_process(seq)
        P = [0]*len(T)
        c,r = 0,0
        for i in range(1,len(T)):
            i_mirror = 2*c - i
            if r > i:
                P[i] = min(r-i, P[i_mirror])
            else:
                P[i] = 0
            while i+1+P[i] < len(T)-1 and i-1-P[i] >=0 and T[ i+1+P[i] ] == T[ i-1-P[i] ]:#比较T[i]左右两边是否相等
                P[i] += 1
            if i + P[i] > r:
                c = i
                r = i+P[i]
        return max(P),P.index(max(P)),self.bef_process(max(P),P.index(max(P)),T),T,P


print Solution().longestPalindrome(s)
