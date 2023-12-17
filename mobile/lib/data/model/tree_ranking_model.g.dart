// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'tree_ranking_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$TreeRankingModelImpl _$$TreeRankingModelImplFromJson(
        Map<String, dynamic> json) =>
    _$TreeRankingModelImpl(
      treeId: json['treeId'] as String,
      rank: json['rank'] as int,
      score: (json['score'] as num).toDouble(),
    );

Map<String, dynamic> _$$TreeRankingModelImplToJson(
        _$TreeRankingModelImpl instance) =>
    <String, dynamic>{
      'treeId': instance.treeId,
      'rank': instance.rank,
      'score': instance.score,
    };
