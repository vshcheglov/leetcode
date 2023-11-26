<?php

use PHPUnit\Framework\TestCase;

class SearchInsertPositionTest extends TestCase
{
    public function test1()
    {
        $solution = new \LeetCode\SearchInsertPosition();
        $result = $solution->searchInsert([1,3,5,6], 5);
        $this->assertEquals(2, $result);
    }

    public function test2()
    {
        $solution = new \LeetCode\SearchInsertPosition();
        $result = $solution->searchInsert([1,3,5,6], 2);
        $this->assertEquals(1, $result);
    }

    public function test3()
    {
        $solution = new \LeetCode\SearchInsertPosition();
        $result = $solution->searchInsert([1,3,5,6], 7);
        $this->assertEquals(4, $result);
    }
}