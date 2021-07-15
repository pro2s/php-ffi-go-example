<?php

namespace FFITest;

require 'vendor/autoload.php';

$product = new Model(1);

$goRunner = new Runner(new GoCombine($product, __DIR__ . "/go/lib.so"));
$phpRunner = new Runner(new Combine($product));
$nodeRunner = new NodeRunner('./nodejs/wasm.js', $product);

$runners = [$phpRunner, $goRunner, $nodeRunner];
array_walk($runners, function (RunnerInterface $runner) {
    $runner->run();
    $runner->print();
});

(new Validator($phpRunner, [$goRunner, $nodeRunner]))->validate();
