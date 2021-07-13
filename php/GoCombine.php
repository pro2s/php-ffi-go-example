<?php

namespace FFITest;

class GoCombine implements CombineInterface
{
    protected $model;
    protected $ffi;
    private const DEFF = <<<C
        typedef struct { char* p; long n } GoString;
        typedef struct { long id; long linked_id; bool empty; } Option;
        typedef struct { void *data; long len; long cap; } GoSlice;
        char* combine(long id, GoSlice data);
    C;

    public function __construct(Model $model, string $lib)
    {
        $this->model = $model;
        $this->ffi = \FFI::cdef(self::DEFF, $lib);
    }

    public function getName(): string
    {
        return 'GO';
    }

    public function combine($combinations)
    {
        $raw = $this->ffi->combine($this->model->id, $combinations);
        $data = \FFI::string($raw);

        return unserialize($data);
    }

    private function newArray($type, $len)
    {
        return $this->ffi->new(\FFI::arrayType($this->ffi->type($type), [$len]), false);
    }

    private function getPtr($data)
    {
        return \FFI::cast(\FFI::type('void *'), $data);
    }

    private function fillGoSlice(\FFI\CData $slice, \FFI\CData $data, int $length)
    {
        $slice->data = $this->getPtr($data);
        $slice->len = $length;
        $slice->cap = $length;
    }

    private function fillOption(\FFI\CData $data, ?Option $option)
    {
        if ($option === null) {
            $data->empty = true;
            return;
        }

        $data->id = $option->id;
        $data->linked_id = $option->linked_id ?? 0;
    }

    public function map(array $combinations): \FFI\CData
    {
        $dataLen = count($combinations);
        $data = $this->newArray('GoSlice', $dataLen);

        foreach ($combinations as $i => $options) {
            $len = count($options);
            $arr = $this->newArray('Option', $len);
            foreach ($options as $j => $option) {
                $this->fillOption($arr[$j], $option);
            }

            $this->fillGoSlice($data[$i], $arr, $len);
        }

        $out = $this->ffi->new('GoSlice', false);
        $this->fillGoSlice($out, $data, $dataLen);

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

    public function formatCombination(array $combinations, $hash): string
    {
        return implode(',', $combinations[$hash] ?? []);
    }
}
