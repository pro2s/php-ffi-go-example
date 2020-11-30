<?php

namespace FFITest;

class Option
{
    public $id;
    public $linked_id;
    public function __construct($id, $linked_id = null)
    {
        [$this->id, $this->linked_id] = [$id, $linked_id];
    }
}
