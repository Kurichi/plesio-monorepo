import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:kiikuten/domain/entity/tree_ranking.dart';

part 'tree_ranking_model.freezed.dart';

part 'tree_ranking_model.g.dart';

@freezed
class TreeRankingModel with _$TreeRankingModel {
  const TreeRankingModel._();

  const factory TreeRankingModel({
    required String treeId,
    required int rank,
    required double score,
  }) = _TreeRankingModel;

  factory TreeRankingModel.fromJson(Map<String, dynamic> json) =>
      _$TreeRankingModelFromJson(json);

  TreeRanking toEntity() {
    return TreeRanking(
      treeId: treeId,
      rank: rank,
      score: score,
    );
  }
}
