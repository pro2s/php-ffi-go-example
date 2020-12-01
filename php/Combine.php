<?php

namespace FFITest;

use FFITest\Model;
use Illuminate\Support\Collection;

class Combine
{
    public $model;

    public function __construct(Model $model)
    {
        $this->model = $model;
    }

    public function getCombinationsArray(array $arrays): array
    {
        if (count($arrays) === 0) {
            return [[]];
        }

        if (count($arrays) === 1) {
            return array_map(fn ($value) => [$value], $arrays[0]);
        }

        $combination = array_shift($arrays);

        $combinations = $this->getCombinationsArray($arrays);

        $result = [];

        foreach ($combination as $value) {
            foreach ($combinations as $combine) {
                $result[] = [$value, ...$combine];
            }
        }

        return $result;
    }

    public function filterLinked($combinations): array
    {
        return array_filter($combinations, static function ($combination) {
            $linked = [[], []];
            foreach ($combination as $option) {
                if ($option) {
                    if ($option->linked_id) {
                        $linked[0][] = (int) $option->linked_id;
                    }
                    $linked[1][] = (int) $option->id;
                }
            }

            return !$linked[0] || count(array_diff(...$linked)) === 0;
        });
    }

    public function getHash(Collection $optionIds): string
    {
        $data = $this->model->id . '-' . $optionIds->implode('id', '-');

        return sprintf("%x", crc32($data));
    }

    public function formatCombinations(array $combinations): array
    {
        $combinationsArray = array_map(function ($combination) {
            $items = collect($combination)->filter();

            return [
                'hash' => $this->getHash($items),
                'ids' => $items->pluck('id')->all(),
                'items' => $items
            ];
        }, $combinations);

        return $combinationsArray;
    }
}
