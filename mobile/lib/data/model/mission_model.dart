import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:kiikuten/data/model/item_model.dart';

part 'mission_model.freezed.dart';

part 'mission_model.g.dart';

@freezed
class MissionModel with _$MissionModel {
  const factory MissionModel({
    required String id,
    required String description,
    required String term,
    required String target,
    required int amount,
    required String unit,
    required List<ItemModel> rewards,
  }) = _MissionModel;

  factory MissionModel.fromJson(Map<String, dynamic> json) =>
      _$MissionModelFromJson(json);
}
