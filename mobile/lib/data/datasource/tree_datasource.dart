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

  Future<void> growTree(String treeId, int growthAmount) async {
    final response = await http.put(
      Uri.parse('$baseUrl/trees/$treeId/grow'),
      body: jsonEncode(<String, int>{'growthAmount': growthAmount}),
      headers: {'Content-Type': 'application/json'},
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to grow tree');
    }
  }
}
