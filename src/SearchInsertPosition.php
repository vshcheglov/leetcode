<?php

namespace LeetCode;

class SearchInsertPosition
{
    function searchInsert($nums, $target) {
        $leftIndex = 0;
        $rightIndex = count($nums) - 1;
        while ($leftIndex <= $rightIndex) {
            $midIndex = floor(($leftIndex + $rightIndex) / 2);
            $midValue = $nums[$midIndex];
            if ($midValue === $target) {
                return $midIndex;
            }
            if ($midValue < $target) {
                $leftIndex = $midIndex + 1;
            } else {
                $rightIndex = $midIndex - 1;
            }
        }
        return $leftIndex;
    }
}
