<?php

namespace FFITest;

class GoCombine
{
    protected $model;
    protected $ffi;

    public function __construct($model, $lib)
    {
        $this->model = $model;
        $this->ffi = \FFI::cdef(
            "
            typedef struct { char* p; long n } GoString;
            typedef struct { long id; long linked_id; bool empty; } Option;
            typedef struct { void *data; long len; long cap; } GoSlice;
            char* combine(long id, GoSlice data);",
            $lib
        );
    }

    public function combine($combinations)
    {
        $raw = $this->ffi->combine($this->model->id, $combinations);
        $data = \FFI::string($raw);

        return unserialize($data);
    }

    public function optionsToGoSlice(array $combinations): \FFI\CData
    {
        $dataLen = count($combinations);
        $data = $this->ffi->new('GoSlice[' . $dataLen . ']', false);

        foreach ($combinations as $di => $options) {
            $len = count($options);
            $arr = $this->ffi->new('Option[' . $len . ']', false);
            foreach ($options as $i => $option) {
                if ($option === null) {
                    $arr[$i]->empty = true;
                    continue;
                }

                $arr[$i]->id = $option->id;
                $arr[$i]->linked_id = $option->linked_id ?? 0;
            }

            $data[$di]->data = \FFI::cast(\FFI::type('void *'), $arr);
            $data[$di]->len = $len;
            $data[$di]->cap = $len;
        }

        $out = $this->ffi->new('GoSlice', false);
        $out->data = \FFI::cast(\FFI::type('void *'), $data);
        $out->len = $dataLen;
        $out->cap = $dataLen;

        return $out;
    }

    public function stringToGoString(string $string): \FFI\CData
    {
        $length = strlen($string);

        $c = \FFI::new('char[' . $length . ']', false);
        \FFI::memcpy($c, $string, strlen($string));

        $goString = $this->ffi->new("GoString");
        $goString->p = \FFI::cast(\FFI::type('char *'), $c);
        $goString->n = $length;

        return $goString;
    }
}
