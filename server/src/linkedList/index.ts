type LlNode = {
  data: any;
  next: LlNode;
};
class LinkedList {
  head: LlNode;
  size: number;
  constructor() {
    this.size = 0;
  }
  add(node: LlNode) {
    if (!this.head) {
      this.head = node;
      this.size = 1;
      return;
    }
    let head = this.head;
    let i = 0;
    while (head.next !== null || i > this.size) {
      head = head.next;
      i++;
    }
    head.next = node;
    this.size++;
  }
  remove(data: any) {
    if (!this.head || this.head.data === data) {
      this.head = this.head.next;
      this.size--;

      return this.head;
    }
    let prev = this.head;
    let next = this.head.next;

    while (prev !== null) {
      prev = this.head;
      next = this.head.next.next;
      if (prev.next.data === data) {
        prev.next = next;
        this.size--;
        break;
      }
    }
  }
  search(data: any) {
    if (this.head && this.head.data == data) {
      console.log('asfdasdsaf', this.head);
      return this.head.data;
    }

    let head = this.head.next;
    console.log('seach');

    while (head != null) {
      if (head.data === data) {
        return head.data;
      }
      head = head.next;
    }

    return null;
  }
}

export function printList() {
  const a: LinkedList = new LinkedList();
  const node: LlNode = { data: 1, next: null };
  a.add(node);
  const node2: LlNode = { data: 2, next: null };
  a.add(node2);
  console.log('linked list search: ', a.search(2));
  a.remove(2);
  console.log(a);
}
