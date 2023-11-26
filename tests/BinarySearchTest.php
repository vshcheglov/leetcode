<?php

use PHPUnit\Framework\TestCase;

class BinarySearchTest extends TestCase
{
    public function test1()
    {
        $solution = new \LeetCode\BinarySearch();
        $result = $solution->search([-1,0,3,5,9,12], 9);
        $this->assertEquals(4, $result);
    }

    public function test2()
    {
        $solution = new \LeetCode\BinarySearch();
        $result = $solution->search([-1,0,3,5,9,12], 2);
        $this->assertEquals(-1, $result);
    }
}