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

    public function getCombinationsArray(array $arrays, int $depth = 0): array
    {
        if (!isset($arrays[$depth])) {
            return [];
        }

        if ($depth === count($arrays) - 1) {
            return $arrays[$depth];
        }

        // get combinations from subsequent arrays
        $combinations = $this->getCombinationsArray($arrays, $depth + 1);

        $result = [];

        // concat each array from combinations with each element from $arrays[$depth]
        foreach ($arrays[$depth] as $value) {
            foreach ($combinations as $combination) {
                $result[] = $this->combine($value, $combination);
            }
        }

        return $result;
    }

    public function combine($value, $combination): array
    {
        if (is_array($combination)) {
            $combination[] = $value;

            return $combination;
        }

        return [$value, $combination];
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
        return hash('crc32', $this->model->id . '-' . $optionIds->implode('id', '-'));
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
