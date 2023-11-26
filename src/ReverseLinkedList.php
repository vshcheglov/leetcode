<?php


namespace LeetCode;

class ReverseLinkedList
{
    function reverseList($head)
    {
        if (!$head) {
            return $head;
        }
        $prev = null;
        $current = $head;
        do {
            $next = $current->next;
            $current->next = $prev;
            $prev = $current;
            $current = $next;
        } while ($next !== null);
        return $prev;
    }
}
