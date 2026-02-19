# 给定索引i,左子节点的索引为2i+1,右子节点的索引为2i+2,父节点的索引为(i-1)//2


class Heap:
    """大顶堆实现"""

    def __init__(self, nums: list[int] = []) -> None:
        self.max_heap = nums
        for i in range(self.parent(self.size() - 1), -1, -1):
            self.sift_down(i)

    def left(self, i: int):
        return 2 * i + 1

    def right(self, i: int):
        return 2 * i + 2

    def parent(self, i: int):
        return (i - 1) // 2

    def peek(self):
        return self.max_heap[0]

    def size(self):
        return len(self.max_heap)

    def is_empty(self):
        return not self.max_heap or self.size() == 0

    def push(self, data: int):
        self.max_heap.append(data)
        self.sift_up(self.size() - 1)

    def pop(self):
        if self.is_empty():
            raise IndexError("heap is empty")
        self.max_heap[0], self.max_heap[self.size() - 1] = (
            self.max_heap[self.size() - 1],
            self.max_heap[0],
        )
        val = self.max_heap.pop()
        self.sift_down(0)

        return val

    def sift_up(self, i):
        while True:
            p = self.parent(i)
            if p < 0 or self.max_heap[p] >= self.max_heap[i]:
                break
            self.max_heap[p], self.max_heap[i] = self.max_heap[i], self.max_heap[p]
            i = p

    def sift_down(self, i):
        while True:
            left, right, max = self.left(i), self.right(i), i
            if left < self.size() and self.max_heap[left] > self.max_heap[i]:
                max = left
            if right < self.size() and self.max_heap[right] > self.max_heap[max]:
                max = right
            if i == max:
                break
            self.max_heap[i], self.max_heap[max] = self.max_heap[max], self.max_heap[i]
            i = max


if __name__ == "__main__":
    heap = Heap([1, 3, 88, 9, 28, 10])

    while not heap.is_empty():
        print(heap.pop())
