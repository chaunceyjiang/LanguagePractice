import xlrd
import xlwt
import os
from collections import defaultdict
from itertools import chain
from datetime import datetime
import re
from pprint import pprint
"""
输入考勤时间表。输出每个人本月对应的迟到（日期）、早退（日期）、缺卡（日期，缺上班卡还是下班卡）、缺工时（日期，缺多少时间）、加班（日期，加班小时数）、请假（日期、时长）等信息。
"""
xlrd.Book.encoding = "utf8"
j = '/home/chauncey/Goexample/pyexample/加班.xls'
s = '/home/chauncey/Goexample/pyexample/时间表.xlsx'
b = '/home/chauncey/Goexample/pyexample/补卡记录.xls'
q = '/home/chauncey/Goexample/pyexample/请假.xls'
if os.path.exists('上班统计.xls'):
    os.remove('上班统计.xls')


def read_xls(filename: str, date_table: list, start=0):
    jdata = xlrd.open_workbook(filename)
    table = jdata.sheets()[0]
    i = 0
    for rownum in range(table.nrows):
        i += 1
        if i >= start:
            date_table.append(table.row_values(rownum))


def fliter_table(table: list, raw_table: dict):
    name = 8
    address = 9
    agree = 3
    buka_time = 14
    pattern = re.compile(r"成都研发")
    buka_name_table = defaultdict(list)
    for l in table:
        if re.match(pattern, l[address]):
            if l[name] in raw_table:
                b_time = [y for y in [x.split()
                                      for x in l[buka_time].split(',')]]
                t = list(chain.from_iterable(b_time))
                t = datetime.strptime(
                    t[0] + " " + t[-1], '%Y-%m-%d %H:%M')
                if l[agree] == "同意":
                    buka_name_table[l[name]].append(t)
    return buka_name_table


def buka_xls(table, raw_table):
    for _, p in enumerate(raw_table):
        # day = "2019-%s-%s %s"
        y_start = 2018
        d_start = 21
        m_start = 12
        flag = False

        # l = len(p_)
        for i in range(5, len(p)):
            if i >= 16 and flag == False:
                y_start = 2019
                m_start = 1
                d_start = 1
                flag = True
            if p[i] != "":
                p_time_sort = sorted([t.strip() for t in p[i].split("\n")])

                work_start = "%s-%s-%s %s" % (y_start, m_start,
                                              d_start, p_time_sort[0])
                work_end = "%s-%0s-%s %s" % (y_start, m_start,
                                             d_start, p_time_sort[-1])
                table[p[0]].append([datetime.strptime(
                    work_start, '%Y-%m-%d %H:%M'), datetime.strptime(work_end, '%Y-%m-%d %H:%M')])
            else:
                table[p[0]].append([None, None])
            d_start += 1


def jiaban_xls(table, raw_table):
    agree = 3
    name = 8
    start = 14
    end = 15
    f = 18
    jiaban_name_table = defaultdict(list)
    for l in table:
        if l[name] in raw_table and l[agree] == "同意":
            jiaban_name_table[l[name]].append([datetime.strptime(
                l[start], '%Y-%m-%d %H:%M'), datetime.strptime(l[end], '%Y-%m-%d %H:%M'), l[f]])
    return jiaban_name_table


def qingjia_xls(table, raw_table):
    agree = 3
    start = 15
    end = 16
    name = 8
    qingjia_name_table = defaultdict(list)
    for l in table:
        if l[name] in raw_table and l[agree] == "同意":
            qingjia_name_table[l[name]].append([datetime.strptime(
                l[start], '%Y-%m-%d %H:%M'), datetime.strptime(l[end], '%Y-%m-%d %H:%M')])
    return qingjia_name_table


date_table = []  # 时间表的整个xls表格
buka_table = []  # 补卡的整个表格
jiaban_table = []  # 加班的整个表格
qiajia_table = []  # 请假的整个表格
read_xls(s, date_table, 4)
read_xls(b, buka_table, 2)
read_xls(j, jiaban_table, 2)
read_xls(q, qiajia_table, 2)

p_time_sort_table = defaultdict(list)  # 保存任务的上下班时间，用来计算
buka_xls(p_time_sort_table, date_table)

buka_name_table = fliter_table(buka_table, p_time_sort_table)  # 补卡的表格
jiaban_name_table = jiaban_xls(jiaban_table, p_time_sort_table)
qingjia_name_table = qingjia_xls(qiajia_table, p_time_sort_table)

save_table_list = []  # 每个sheet表单对象
pprint(p_time_sort_table)


def write_xml(sava_tanle_list: list, filename='上班统计.xls',):
    file = xlwt.Workbook()
    m_start_time = datetime.strptime("2018-12-21 9:00", '%Y-%m-%d %H:%M')
    d_start_time = "{}-{}-{} 9:30"
    d_mid_time = "{}-{}-{} 13:30"
    d_end_time_1 = "{}-{}-{} 18:00"
    d_end_time_2 = "{}-{}-{} 18:30"
    name_list = ['迟到', '早退', '缺卡', '缺工时', '加班', '请假']
    value_list = [0] * 6
    for i, v in enumerate(date_table):
        save_table = file.add_sheet(v[0], cell_overwrite_ok=True)
        for t in p_time_sort_table[v[0]]:
            if t[0] is None and t[1] is None:
                continue
            year, month, day = t[0].year, t[0].month, t[0].day
            if  t[0] > datetime.strptime(  # 迟到
                    d_start_time.format(
                        year,
                        month,
                        day),
                    '%Y-%m-%d %H:%M'):
                value_list[0] += 1
            if t[1] < datetime.strptime(    # 早退
                    d_end_time_1.format(
                        year,
                        month,
                        day),
                    '%Y-%m-%d %H:%M'):
                value_list[1] += 1
            elif t[0]<=datetime.strptime(  # 早退
                    d_start_time.format(
                        year,
                        month,
                        day),
                    '%Y-%m-%d %H:%M') and t[1]<datetime.strptime(
                    d_end_time_2.format(
                        year,
                        month,
                        day),
                    '%Y-%m-%d %H:%M') and (t[1] - t[0]).seconds <= 32400:
                value_list[1] += 1

            if  (t[1] - t[0]).seconds <= 32400:    # 缺工时
                value_list[3] += 1



        # save_table.write(0,0,"A")
        sava_tanle_list.append(save_table)
        for i in range(6):
            save_table.write(0, i, name_list[i])

    file.save(filename)


write_xml(save_table_list)


# os.remove('上班统计.xls')
