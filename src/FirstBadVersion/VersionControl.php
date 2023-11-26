<?php

namespace LeetCode\FirstBadVersion;

class VersionControl
{
    private $firstBad;

    public function __construct(int $firstBad)
    {
        $this->firstBad = $firstBad;
    }

    public function isBadVersion(int $version)
    {
        return $version >= $this->firstBad;
    }
}
