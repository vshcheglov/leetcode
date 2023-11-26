<?php

namespace LeetCode;

class BinarySearch
{
    function search($nums, $target)
    {
        $leftIndex = 0;
        $rightIndex = count($nums) - 1;
        while ($leftIndex <= $rightIndex) {
            $middleIndex = floor(($leftIndex + $rightIndex) / 2);
            $middleValue = $nums[$middleIndex];
            if ($middleValue == $target) {
                return $middleIndex;
            }
            if ($middleValue < $target) {
                $leftIndex = $middleIndex + 1;
            } else {
                $rightIndex = $middleIndex - 1;
            }
        }
        return -1;
    }
}
