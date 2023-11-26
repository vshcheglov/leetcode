<?php

use LeetCode\ReverseLinkedList\ListNode;
use PHPUnit\Framework\TestCase;

class ReverseLinkedListTest extends TestCase
{
    public function test1()
    {
        $solution = new \LeetCode\ReverseLinkedList();
        $node = ListNode::fromArray([1,2,3,4,5]);
        $result = $solution->reverseList($node);
        $this->assertEquals([5,4,3,2,1], ListNode::toArray($result));
    }

    public function test2()
    {
        $solution = new \LeetCode\ReverseLinkedList();
        $node = ListNode::fromArray([1,2]);
        $result = $solution->reverseList($node);
        $this->assertEquals([2,1], ListNode::toArray($result));
    }

    public function test3()
    {
        $solution = new \LeetCode\ReverseLinkedList();
        $node = ListNode::fromArray([]);
        $result = $solution->reverseList($node);
        $this->assertEquals([], ListNode::toArray($result));
    }
}
