<?php

use PHPUnit\Framework\TestCase;

class TwoSumTest extends TestCase
{
    public function test1()
    {
        $solution = new \LeetCode\TwoSum();
        $result = $solution->twoSum([2,7,11,15], 9);
        $this->assertEquals([0,1], $result);
    }

    public function test2()
    {
        $solution = new \LeetCode\TwoSum();
        $result = $solution->twoSum([3,2,4], 6);
        $this->assertEquals([1,2], $result);
    }

    public function test3()
    {
        $solution = new \LeetCode\TwoSum();
        $result = $solution->twoSum([3,3], 6);
        $this->assertEquals([0,1], $result);
    }
}
