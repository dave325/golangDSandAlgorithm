import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';
import { printList as linkedListPrintList } from './linkedList';
import { printList as stackPrintList } from './stack';
import { printList as heapPrintList } from './heap';
@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('linkedList')
  lnkedList(): string {
    linkedListPrintList();
    return this.appService.getHello();
  }

  @Get('stack')
  stack(): string {
    stackPrintList();
    return this.appService.getHello();
  }

  @Get('heap')
  heap(): string {
    heapPrintList();
    return this.appService.getHello();
  }
}
