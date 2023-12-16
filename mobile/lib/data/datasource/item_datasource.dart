import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:kiikuten/data/model/item_model.dart';

class ItemDataSource {
  final String baseUrl;

  ItemDataSource(this.baseUrl);

  Future<ItemModel> getItem(String itemId) async {
    final response = await http.get(Uri.parse('$baseUrl/items/$itemId'));

    if (response.statusCode == 200) {
      return ItemModel.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load item');
    }
  }
}
