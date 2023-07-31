import queue

# 创建一个先进先出队列
q = queue.Queue()

# 入队
q.put(1)
q.put(2)
q.put(3)

# 出队
while not q.empty():
    item = q.get()
    print(item)