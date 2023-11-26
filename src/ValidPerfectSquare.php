<?php

namespace LeetCode;

class ValidPerfectSquare
{
    public function isPerfectSquare($num)
    {
        $left = 0;
        $right = ceil($num / 2);
        while ($left <= $right) {
            $mid = floor(($left + $right) / 2);
            if ($mid * $mid == $num) {
                return true;
            }
            if ($mid * $mid < $num) {
                $left = $mid + 1;
            } else {
                $right = $mid - 1;
            }
        }
        return false;
    }
}
