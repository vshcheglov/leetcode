<?php

use PHPUnit\Framework\TestCase;

class AddTwoIntegersTest extends TestCase
{
    public function test1()
    {
        $solution = new \LeetCode\AddTwoIntegers();
        $result = $solution->sum(12, 5);
        $this->assertEquals(17, $result);
    }

    public function test2()
    {
        $solution = new \LeetCode\AddTwoIntegers();
        $result = $solution->sum(-10, 4);
        $this->assertEquals(-6, $result);
    }
}
