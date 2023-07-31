class Queue {
  constructor() {
    this.items = [];
  }

  enqueue(item) {
    this.items.push(item);
  }

  dequeue() {
    if (this.isEmpty()) {
      return -1;
    }
    return this.items.shift();
  }

  isEmpty() {
    return this.items.length === 0;
  }
}

const queue = new Queue();
queue.enqueue(1);
queue.enqueue(2);
queue.enqueue(3);

console.log(queue.dequeue()); // 输出：1
console.log(queue.dequeue()); // 输出：2
console.log(queue.dequeue()); // 输出：3
console.log(queue.dequeue()); // 输出：-1（队列为空）