-- 关于 LIMIT 的用法
-- select * from Customer LIMIT 10;--检索前10行数据，显示1-10条数据
-- select * from Customer LIMIT 1,10;--检索从第2行开始，累加10条id记录，共显示id为2....11
-- select * from Customer limit 5,10;--检索从第6行开始向前加10条数据，共显示id为6,7....15
-- select * from Customer limit 6,10;--检索从第7行开始向前加10条记录，显示id为7,8...16

WHERE 和HAVING： 
WHERE语句在GROUP BY之前；SQL会在分组前计算WHERE语句 
HAVING语句在GROUP BY之后；SQL会在分组后计算HAVING语句

--limit用法
SELECT * FROM tablename LIMIT 0,1
即取出第一条记录。

SELECT * FROM tablename LIMIT 1,1
第二条记录

SELECT * FROM tablename LIMIT 10,20
从第11条到31条（共计20条）

--------------------------------------------------------------------------------------


SELECT * FROM xxx ORDER BY hire_datea DESC LIMIT 2,1; -- 查找入职员工时间排名倒数第三的员工所有信息 

SELECT DISTINCT country FROM Websites;
-- 仅从 "Websites" 表的 "country" 列中选取唯一不同的值，也就是去掉 "country" 列重复值

SELECT * from websites where country='cn';  --大小写不区分，cn 与 CN 都可以。
SELECT * from websites where country='cn' and alexa >= 20;  


SELECT * FROM websites WHERE country='CN' AND alexa > 50;


SELECT * FROM websites WHERE country='USA' OR country='CN';


SELECT * FROM websites WHERE alexa > 15 AND (country='CN' OR country='USA');

SELECT * FROM websites WHERE alexa > 15 AND (country='CN' OR country='USA') ORDER BY alexa desc; --desc降序 ASC升序


--  逻辑运算的优先级：
--         ()    not        and         or


--  特殊条件

-- 1.空值判断： is null

--         Select * from websites where name is null;

-- 2.between and (在 之间的值)

-- 3.In


-- ORDER BY 多列的时候，先按照第一个column name排序，在按照第二个column name排序；如上述教程最后一个例子：
--     1）、先将country值这一列排序，同为CN的排前面，同属USA的排后面；
--     2）、然后在同属CN的这些多行数据中，再根据alexa值的大小排列。
--     3）、ORDER BY 排列时，不写明ASC DESC的时候，默认是ASC。 