import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class Leetcode {
    public static void main(String[] args) {
        Solution s = new Solution();
        String ss = "abcbaccax";
//        String.join()
        int[] result;
        int[] ax = {1, 4, 9, 34, 2, 7, 94, 32, 55, 62, 3, 45, 24, 36, 8, 4, 74, 48};
        int[] ay = {3, 3, 3, 3, 3, 3, 3, 3, 3, 3};
        int[] az = {2, 1,2,4,5,4,5};
//        s.threeSum(ax);
//        ArrayList<Integer> l = new ArrayList<Integer>(Arrays.asList(ax));
//        System.out.println(s.threeSum(ax));
        s.reverse(2147483647);
        ListNode root = new ListNode(1);
        ListNode t = root;
        for (int i = 2; i < 5; i++) {
            t.next = new ListNode(i + i);

            t = t.next;
        }

//        root=s.addTwoNumbers(root,root);

        System.out.println(ax.length / 2);
        String[] sx;
        Set<Integer> set = new HashSet<>();
        System.out.println(set.add(1));
        System.out.println(set.add(1));
        sx = ss.split("");
        ss = String.join("#", sx);
        System.out.println(ss);
        System.out.println('1' - 'a');
        System.out.println(3 >> 1);
        System.out.println(5^2^6^2^6);
        String x = Integer.toString(123445);
        String xx = new StringBuffer(x).reverse().toString();
        int st =Integer.valueOf("2324");
    }
}
class Solution {
    public int reverse(int x) {
        int rev = 0;
        while (x!=0){
            int pop = x%10;
            x/=10;
            if(rev>Integer.MAX_VALUE/10||(rev==Integer.MAX_VALUE/10 && pop>7))return 0;
            if(rev<Integer.MIN_VALUE/10||(rev==Integer.MIN_VALUE/10 && pop<-8))return 0;

            rev=rev*10 + pop;
        }
        return rev;
    }
}
//class Solution {
//    public int singleNumber(int[] nums) {
//        int result =0;
//        Map<Integer,Integer> m = new HashMap<>();
//        for (int i:nums){
//            m.put(i,m.getOrDefault(i,0)+1);
//        }
//        for(int k:m.keySet()){
//            if (m.get(k)==1){
//                result=k;
//                break;
//            }
//        }
//
//        return result;
//    }
//}
//class Solution {
//    private boolean exit = false;
//    private int value;
//
//    public int findKthLargest(int[] nums, int k) {
//        QuickSort(nums, 0, nums.length - 1, nums.length - k);
//        if (exit) {
//            return value;
//        }
//        return nums[nums.length - k];
//    }
//
//    private void QuickSort(int[] nums, int left, int right, int k) {
//        if (exit) {
//            return;
//        }
//        if (left < right) {
//            int i = partition(nums, left, right, k);
//            QuickSort(nums, left, i-1, k);
//            QuickSort(nums, i + 1, right, k);
//        }
//    }
//
//    private int partition(int[] nums, int left, int right, int k) {
//        if (exit) {
//            return value;
//        }
//        int pivot = nums[left];
//        while (left<right){
//            while (left<right&&nums[right]>=pivot) right--;
//            nums[left]=nums[right];
//            while (left<right&&nums[left]<=pivot) left++;
//            nums[right]=nums[left];
//        }
//        nums[left] =pivot;
//        if (left==k){
//            exit=true;
//            value=nums[k];
//        }
//        return left;
//    }
//}

//class Solution {
//    int max =0;
//    public int maxDepth(TreeNode root) {
//        dfs(root,0);
//        return max;
//
//    }
//    private void dfs(TreeNode root,int h){
//        if(root==null){
//            return;
//        }
//        if(h+1>max){
//            max=h+1;
//        }
//        dfs(root.left,h+1);
//        dfs(root.right,h+1);
//    }
//}
//class Solution {
//    public int majorityElement(int[] nums) {
//        Map<Integer,Integer> map = new HashMap<>();
//        for (int i=0;i<nums.length;i++){
//            map.put(nums[i],map.getOrDefault(nums[i],0)+1);
//        }
//        int max =0;
//        int res=0;
//        for (int i:map.keySet()){
//            if(map.get(i)>max){
//                res = i;
//                max=map.get(i);
//            }
//        }
//        return res;
//    }
//}

//class Solution {
//    public boolean hasCycle(ListNode head) {
//        ListNode slow ,fast;
//        if (head==null || head.next==null||head.next.next==null){
//            return false;
//        }
//        slow=head;
//        fast=head.next;
//        while (fast!=slow){
//            if (fast == null || fast.next == null){
//                return false;
//            }
//            fast = fast.next.next;
//            slow = slow.next;
//        }
//        return true;
//    }
//}

//class Solution {
//    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
//        return add(l1,l2,0);
//    }
//    private ListNode add(ListNode l1,ListNode l2 ,int p){
//
//        if (l1 == null && l2 == null){
//            if (p==0)
//                return null;
//            else
//                return new ListNode(p);
//        }
//        int x,y,sum;
//        ListNode t1 ,t2;
//        boolean f = false;
//
//        if (l1 == null){
//            x=0;
//            t1 = null;
//        }else {
//            x = l1.val;
//            t1 = l1.next;
//        }
//        if (l2 == null){
//            y=0;
//            t2 = null;
//        }else {
//            y = l2.val;
//            t2 = l2.next;
//        }
//
//        sum = x +y+p;
//
//        if (sum>=10){
//            f = true;
//            sum %=10;
//        }
//        ListNode t =new ListNode(sum);
//        t.next=add(t1,t2,f?1:0);
//        return t;
//    }
//}

//class Solution {
//    public ListNode reverseList(ListNode head) {
//
//        ListNode p =head;
//        if (head == null){
//            return null;
//        }
//        if (head.next==null){
//            return p;
//        }
//        if (head.next.next == null){
//            ListNode cur = head.next;
//            cur.next = p;
//            p.next = null;
//            return cur;
//        }
//        ListNode cur = head.next;
//        head.next=null;
//        ListNode res = null;
//
//        while (cur!=null){
//            res = cur.next;
//            cur.next = p;
//            p = cur;
//            cur = res;
//        }
//        return p;
//    }
//
//}

//class Solution {
//    public List<List<Integer>> threeSum(int[] nums) {
//        List<List<Integer>> l = new ArrayList<>();
//        Arrays.sort(nums);
//        for(int i=0;i<nums.length-2;i++){
//            // 去重
//            if(i>0 && nums[i]==nums[i-1]){
//                continue;
//            }
//            int low=i+1,high=nums.length-1,sum=0-nums[i];
//            while(low<high){
//                //如果符合条件
//                if(nums[low]+nums[high]==sum){
//                    l.add(Arrays.asList(nums[i],nums[low],nums[high]));
//                    //去重
//                    while(low<high && nums[low]==nums[low+1]) low++;
//                    while(low<high && nums[high]==nums[high-1]) high--;
//                    low++;
//                    high--;
//                }
//                //如果小了，把小的右移一位
//                else if (nums[low]+nums[high]<sum){
//                    low++;
//                }
//                //如果大了，把大的左移一位
//                else {
//                    high--;
//                }
//            }
//        }
//        return l;
//    }
//}

//class Solution {
//    public List<List<Integer>> threeSum(int[] nums) {
//        Arrays.sort(nums);
//        List<List<Integer>> l = new ArrayList<>();
//        Set<List<Integer>> set = new HashSet<>();
//        if (nums.length<3){
//            return l;
//        }
//
//        for (int i=0;i<nums.length;i++){
//            for (int j=i+1;j<nums.length;j++){
//                for (int k=j+1;k<nums.length;k++){
//
//                    List<Integer> tmpl = new ArrayList<>();
//                    int[] n = new int[3];
//                    n[0]= nums[i];
//                    n[1]= nums[j];
//                    n[2]= nums[k];
//                    Arrays.sort(n);
//                    tmpl.add(n[0]);
//                    tmpl.add(n[1]);
//                    tmpl.add(n[2]);
//
//                    if (nums[i]+nums[j]+nums[k]==0 && set.add(tmpl)){
//                        l.add(tmpl);
//                    }
//                }
//            }
//        }
//        return l;
//    }
//
//
//}

//class Solution {
//    public String longestPalindrome(String s) {
//        int n = s.length();
//        if (n==0){
//            return s;
//        }
//        boolean[][] pal = new boolean[n][n]; //用来记录i到j是否是回文
//        int maxLen = 0,start=0,end=0;
//        for(int i=0;i<n;i++){
//
//            int j=i; //每次都向前检查是否是回文
//
//            while (j>=0){
//                if(s.charAt(j)==s.charAt(i)&&((i-j<2) || pal[j+1][i-1])){// 如果当前元素相等，并且它们之间的也都回文，则当前也是回文
//                    pal[j][i]=true;
//                    if (maxLen<=i-j+1){
//                        maxLen=i-j+1;
//                        start=j;
//                        end=i;
//                    }
//
//                }
//                j--;
//            }
//        }
//        return s.substring(start,end+1);
//    }
//}
//class Solution {
//    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
//        int nums2length = nums2.length;
//        int[] r = new int[nums1.length + nums2length];
//
//        int index = 0;
//        int j = 0;
//        for (int num : nums1) {
//            while (j < nums2length && num > nums2[j]) {// 因为是有序的，所以找出当前数组各自最小的元素
//                r[index] = nums2[j];
//                j++;
//                index++;
//            }
//            r[index] = num;
//            index++;
//        }
//        for (; j < nums2length; j++) {// 补充 数组2 后面都大于数组1的元素
//            r[index] = nums2[j];
//            index++;
//        }
//
//        int mi = r.length / 2;
//        return r.length % 2 == 1 ? r[mi] : (r[mi - 1] + r[mi]) / 2.0;
//    }
//}
//class Solution {
//    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
//        if (nums1.length == 0) {
//            return checkNullArray(nums2);
//        }
//        if (nums2.length == 0) {
//            return checkNullArray(nums1);
//        }
//        int[] result;
//        result=concat(nums1,nums2);
//        return checkNullArray(result);
//    }
//
//    private double checkNullArray(int[] a) {
//        if (a.length % 2 == 0) {
//            return (a[a.length / 2] + a[a.length / 2 - 1]) / 2.0;
//
//        }
//        return a[a.length / 2] ;
//
//    }
//    private int[] concat(int[] a ,int[]b){
//        int[] n = new int[a.length+b.length];
//        System.arraycopy(a,0,n,0,a.length);
//        System.arraycopy(b,0,n,a.length,b.length);
//        Arrays.sort(n);
//        return n;
//    }
//}


//class Solution {
//    public int[] twoSum(int[] nums, int target) {
//        int[] result = new int[2];
//        Map<Integer,Integer> m = new HashMap<>();
//        for (int i=0;i<nums.length;i++){
//            if (m.containsKey(nums[i])){
//                result[0]=m.get(nums[i]);
//                result[1]=i;
//                return result;
//            }
//            m.put(target-nums[i],i);
//        }
//
//        return result;
//    }
//}
//class Solution {
//    public int numberOfBoomerangs(int[][] points) {
//        int result = 0;
//        Map<Integer, Integer> map;
//        for (int i = 0; i < points.length; i++) {
//            map = new HashMap<>();
//            // 计算当前点到其他每个点的距离 从相同距离中选择两个点 排列组合 有序An2 从n个中有序的选择2个
//            for (int j = 0; j < points.length; j++) {
//                // 不考虑当前同样的点
//                if (i != j) {
//                    int distance = getDistance(points[i], points[j]);
//                    map.put(distance, map.getOrDefault(distance, 0) + 1);
//                }
//            }
//            result += getCountByPoint(map);
//        }
//
//        return result;
//    }
//
//    /**
//     * 求两点之间的距离的平方, 不进行平方根操作了, 不然可能会涉及到浮点数问题
//     *
//     * @param pointA
//     * @param pointB
//     * @return
//     */
//    private int getDistance(int[] pointA, int[] pointB) {
//        int x = pointA[0] - pointB[0];
//        int y = pointA[1] - pointB[1];
//        ArrayList a = new ArrayList();
//
//        return x * x + y * y;
//    }
//
//    /**
//     * 根据点的数量返回回旋镖的数量
//     *
//     * @return
//     */
//    private int getCountByPoint(Map<Integer, Integer> map) {
//        int result = 0;
//        for (Integer value : map.values()) {
//            if (value > 1) {
//                // 从n个点中选取两个点的方案数量为(n * n - 1)
//                result += value * (value - 1);
//            }
//        }
//        return result;
//    }
//}

//class Solution {
//    public int numJewelsInStones(String J, String S) {
//        int count=0;
//        HashMap<Character,Integer> m = new HashMap<Character, Integer>();
//        for (int j=0;j<J.length();j++){
//            m.put(J.charAt(j),0);
//        }
//        for (int i=0;i<S.length();i++){
//            if(m.get(S.charAt(i))!=null){
//                m.put(S.charAt(i),m.get(S.charAt(i))+1);
//            }
//        }
//
//        for (Character key:m.keySet()){
//            count+=m.get(key);
//        }
//        return count;
//    }
//}