<?php

namespace LeetCode;

use LeetCode\FirstBadVersion\VersionControl;

class FirstBadVersion extends VersionControl
{
    function firstBadVersion($n)
    {
        $leftIndex = 1;
        $rightIndex = $n;
        $firstBadVersion = null;
        while ($leftIndex <= $rightIndex) {
            $midIndex = floor(($leftIndex + $rightIndex) / 2);
            $isMidBadVersion = $this->isBadVersion($midIndex);
            if ($isMidBadVersion) {
                $firstBadVersion = $midIndex;
                $rightIndex = $midIndex - 1;
            } else {
                $leftIndex = $midIndex + 1;
            }
        }
        return $firstBadVersion;
    }
}
