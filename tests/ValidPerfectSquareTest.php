<?php

use PHPUnit\Framework\TestCase;

class ValidPerfectSquareTest extends TestCase
{
    public function test1()
    {
        $solution = new \LeetCode\ValidPerfectSquare();
        $result = $solution->isPerfectSquare(16);
        $this->assertEquals(true, $result);
    }

    public function test2()
    {
        $solution = new \LeetCode\ValidPerfectSquare();
        $result = $solution->isPerfectSquare(14);
        $this->assertEquals(false, $result);
    }

    public function test3()
    {
        $solution = new \LeetCode\ValidPerfectSquare();
        $result = $solution->isPerfectSquare(0);
        $this->assertEquals(true, $result);
    }

    public function test4()
    {
        $solution = new \LeetCode\ValidPerfectSquare();
        $result = $solution->isPerfectSquare(1);
        $this->assertEquals(true, $result);
    }

    public function test5()
    {
        $solution = new \LeetCode\ValidPerfectSquare();
        $result = $solution->isPerfectSquare(2);
        $this->assertEquals(false, $result);
    }

    public function test6()
    {
        $solution = new \LeetCode\ValidPerfectSquare();
        $result = $solution->isPerfectSquare(4);
        $this->assertEquals(true, $result);
    }
}
