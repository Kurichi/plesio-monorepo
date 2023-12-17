import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:kiikuten/data/model/tree_model.dart';
import 'package:kiikuten/data/model/tree_ranking_model.dart';

class TreeDataSource {
  final String baseUrl;

  TreeDataSource(this.baseUrl);

  Future<TreeModel> getMyTree() async {
    final response = await http.get(Uri.parse('$baseUrl/api/v1/trees/me'));

    if (response.statusCode == 200) {
      return TreeModel.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load tree');
    }
  }

  Future<TreeModel> getTreeByUserId(String userId) async {
    final response =
        await http.get(Uri.parse('$baseUrl/api/v1/users/$userId/tree'));

    if (response.statusCode == 200) {
      return TreeModel.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load user tree');
    }
  }

  Future<List<TreeRankingModel>> getTreeRanking() async {
    final response = await http.get(Uri.parse('$baseUrl/api/v1/trees/ranking'));

    if (response.statusCode == 200) {
      Iterable jsonResponse = json.decode(response.body);
      return List<TreeRankingModel>.from(
        jsonResponse.map((model) => TreeRankingModel.fromJson(model)),
      );
    } else {
      throw Exception('Failed to load tree rankings');
    }
  }

  Future<void> plantTree() async {
    final response = await http.post(Uri.parse('$baseUrl/api/v1/trees/plant'));

    if (response.statusCode != 200) {
      throw Exception('Failed to plant tree');
    }
  }

  Future<void> initTree() async {
    final response = await http.post(Uri.parse('$baseUrl/api/v1/trees/init'));

    if (response.statusCode != 200) {
      throw Exception('Failed to initialize tree');
    }
  }
}
