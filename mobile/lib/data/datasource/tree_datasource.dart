import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:kiikuten/data/model/tree_model.dart';

class TreeDataSource {
  final String baseUrl;

  TreeDataSource(this.baseUrl);

  Future<TreeModel> getTree(String treeId) async {
    final response = await http.get(Uri.parse('$baseUrl/trees/$treeId'));

    if (response.statusCode == 200) {
      return TreeModel.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load tree');
    }
  }
}
