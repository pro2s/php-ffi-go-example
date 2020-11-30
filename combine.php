<?php

function stringToGoString(FFI $ffi, string $name): FFI\CData
{
    $strChar = str_split($name);

    $c = FFI::new('char[' . count($strChar) . ']', false);
    foreach ($strChar as $i => $char) {
        $c[$i] = $char;
    }

    $goStr = $ffi->new("GoString");
    $goStr->p = FFI::cast(FFI::type('char *'), $c);
    $goStr->n = count($strChar);

    return $goStr;
}

$ffi = FFI::cdef(
    "typedef struct { char* p; long n } GoString;
    char* combine(GoString data);",
    __DIR__ . "/go/libutil.so"
);
class Id
{
    public $id;
    public $link_id;
}
$url = stringToGoString($ffi, serialize(
    [
        [['id' => 1, 'link_id' => 1],['id' => 2, 'link_id' => null]],
        [['id' => 10, 'link_id' => 1],['id' => 12, 'link_id' => 1]]
    ],
));

$back = FFI::string($ffi->combine($url));
echo $back;
var_dump(unserialize($back));
