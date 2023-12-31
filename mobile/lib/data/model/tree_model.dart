import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:kiikuten/domain/entity/tree.dart';

part 'tree_model.freezed.dart';

part 'tree_model.g.dart';

@freezed
class TreeModel with _$TreeModel {
  const TreeModel._();

  const factory TreeModel({
    required String id,
    required int growthLevel,
  }) = _TreeModel;

  factory TreeModel.fromJson(Map<String, dynamic> json) =>
      _$TreeModelFromJson(json);

  Tree toEntity() {
    return Tree(
      id: id,
      growthLevel: growthLevel,
    );
  }
}
