<?php

use PHPUnit\Framework\TestCase;

class MergeSortedArrayTest extends TestCase
{
    public function test1()
    {
        $solution = new \LeetCode\MergeSortedArray();
        $array = [1,2,3,0,0,0];
        $solution->merge($array, 3,  [2,5,6], 3);
        $this->assertEquals([1,2,2,3,5,6], $array);
    }

    public function test2()
    {
        $solution = new \LeetCode\MergeSortedArray();
        $array = [1];
        $solution->merge($array, 1,  [], 0);
        $this->assertEquals([1], $array);
    }

    public function test3()
    {
        $solution = new \LeetCode\MergeSortedArray();
        $array = [0];
        $solution->merge($array, 0,  [1], 1);
        $this->assertEquals([1], $array);
    }
}
