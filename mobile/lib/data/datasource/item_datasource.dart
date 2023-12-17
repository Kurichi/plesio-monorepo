import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:kiikuten/data/model/item_model.dart';

class ItemDataSource {
  final String baseUrl;

  ItemDataSource(this.baseUrl);

  Future<List<ItemModel>> getItems() async {
    final response = await http
        .get(Uri.parse('$baseUrl/api/v1/items'))
        .timeout(const Duration(seconds: 10));

    if (response.statusCode == 200) {
      List<dynamic> body = json.decode(response.body);
      return body.map((dynamic item) => ItemModel.fromJson(item)).toList();
    } else {
      throw Exception('Failed to load items: ${response.statusCode}');
    }
  }

  Future<void> useItem(String itemId) async {
    final response = await http
        .post(Uri.parse('$baseUrl/api/v1/items/$itemId/use'))
        .timeout(const Duration(seconds: 10));

    if (response.statusCode != 200) {
      throw Exception('Failed to use item: ${response.statusCode}');
    }
  }
}
