type Element = {
  priority: number;
  data: any;
};
class PriorityQueue {
  private Arr: Element[] = [];
  constructor(arr: Element[]) {
    this.Arr = arr;
  }

  max_heapify(i: number, N: number) {
    let largest;
    const left = 2 * i; //left child
    const right = 2 * i + 1; //right child
    if (left <= N && this.Arr[left]?.priority > this.Arr[i]?.priority)
      largest = left;
    else largest = i;
    if (right <= N && this.Arr[right]?.priority > this.Arr[largest]?.priority)
      largest = right;
    if (largest != i) {
      this.swap(i, largest);
      this.max_heapify(largest, N);
    }
  }

  swap(i: number, largest: number) {
    const largestCopy = this.Arr[largest];
    const iCopy = this.Arr[i];
    this.Arr[largest] = iCopy;
    this.Arr[i] = largestCopy;
  }

  build_maxheap(arr: Element[]) {
    this.Arr = arr;
    for (let i = this.Arr.length / 2; i >= 1; i--) {
      this.max_heapify(i, this.Arr.length);
    }
  }

  heap_sort(arr: Element[]) {
    let heap_size = arr.length;
    this.build_maxheap(this.Arr);
    for (let i = arr.length; i >= 2; i--) {
      this.swap(1, i);
      heap_size = heap_size - 1;
      this.max_heapify(1, heap_size);
    }
  }

  maximum() {
    return this.Arr[1]; //as the maximum element is the root element in the max heap.
  }

  extract_maximum() {
    if (this.Arr.length == 0) {
      console.log('Can’t remove element as queue is empty');
      return null;
    }
    const max = this.Arr[1];
    this.Arr[1] = this.Arr[this.Arr.length - 1];
    this.Arr.splice(-1,1);
    
    const length = this.Arr.length - 1;
    this.max_heapify(1, length);
    return max;
  }
}

export function printList() {
  let arr = [null, 6, 4, 5, 3, 2, 0, 1].map((x) => {
    return { priority: x, data: `he - ${x}` };
  });
  const a = new PriorityQueue(arr);

  console.log(a, a.extract_maximum());
  arr = [null, 1, 4, 3, 7, 8, 9, 10].map((x) => {
    return { priority: x, data: `he - ${x}` };
  });
  a.build_maxheap(arr);
  console.log(a, a.extract_maximum());
  arr = [null, 4, 3, 7, 1, 8, 5].map((x) => {
    return { priority: x, data: `he - ${x}` };
  });
  a.heap_sort(arr);
  console.log(a, a.extract_maximum());
}
