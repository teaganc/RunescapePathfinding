#include "ch_map.h"
#include <fstream>
#include <sstream>

void CHMap::PreloadMap(std::string map_file, std::string coord_file) {
    ch = RoutingKit::ContractionHierarchy::load_file(map_file);

    std::ifstream f(coord_file);

	if (f.is_open() && f.good()) {
        std::string line;
		int a, b, c;

        while(std::getline(f, line)) {
			std::stringstream ss(line);
            ss >> a;
            ss.ignore();
            ss >> b;
            ss.ignore();
            ss >> c;
            std::pair<int, int> point = {b, c};
            index_to_point_[a] = point;
            point_to_index_[point] = a;
        }
        f.close();
    }
}

std::vector<std::pair<int, int>> CHMap::GetPath(std::pair<int, int> start, std::pair<int, int> end) {
    if (point_to_index_.contains(start) && point_to_index_.contains(end)) {
        int start_index = point_to_index_.at(start);
        int end_index = point_to_index_.at(end);

        RoutingKit::ContractionHierarchyQuery query(ch);
	    query.reset().add_source(start_index).add_target(end_index).run();

        std::vector<std::pair<int,int>> path;

        for (auto i: query.get_node_path()){
            path.push_back(index_to_point_.at(i));
        }
        return path;
    }
    return {};
}