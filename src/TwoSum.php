<?php

namespace LeetCode;

class TwoSum
{
    function twoSum($nums, $target)
    {
        $map = [];
        foreach($nums as $index => $num) {
            if (isset($map[$target - $num])) {
                return [$map[$target - $num], $index];
            }
            $map[$num] = $index;
        }
        return null;
    }
}
