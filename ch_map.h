#pragma once 

#include <routingkit/contraction_hierarchy.h>
#include <unordered_map>
#include <utility>
#include <vector>


struct pair_hash {
    template<typename T, typename U>
    std::size_t operator() (const std::pair<T, U> &pair) const {
        return std::hash<T>()(pair.first) ^ std::hash<U>()(pair.second);
    }
};

class CHMap {
    RoutingKit::ContractionHierarchy ch; 

    std::unordered_map<int, std::pair<int, int>> index_to_point_;
    std::unordered_map<std::pair<int, int>, int, pair_hash> point_to_index_;

public:
    void PreloadMap(std::string map_file, std::string coord_file);

    std::vector<std::pair<int, int>> GetPath(std::pair<int, int> start, std::pair<int, int> end);

};