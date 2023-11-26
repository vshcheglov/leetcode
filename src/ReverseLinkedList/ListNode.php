<?php

namespace LeetCode\ReverseLinkedList;

class ListNode
{
    public $val = 0;
    public $next = null;

    function __construct($val = 0, $next = null)
    {
        $this->val = $val;
        $this->next = $next;
    }

    public static function toArray(ListNode $node = null)
    {
        $result = [];
        do {
            if ($node) {
                $result[] = $node->val;
                $node = $node->next;
            }
        } while ($node);
        return $result;
    }

    public static function fromArray($array)
    {
        $node = null;
        $next = null;
        foreach (array_reverse($array) as $item) {
            $node = new ListNode($item);
            $node->val = $item;
            $node->next = $next;
            $next = $node;
        }
        return $node;
    }
}
