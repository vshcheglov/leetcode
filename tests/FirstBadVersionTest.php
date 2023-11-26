<?php

use PHPUnit\Framework\TestCase;

class FirstBadVersionTest extends TestCase
{
    public function test1()
    {
        $solution = new \LeetCode\FirstBadVersion(4);
        $result = $solution->firstBadVersion(5);
        $this->assertEquals(4, $result);
    }

    public function test2()
    {
        $solution = new \LeetCode\FirstBadVersion(1);
        $result = $solution->firstBadVersion(1);
        $this->assertEquals(1, $result);
    }
}