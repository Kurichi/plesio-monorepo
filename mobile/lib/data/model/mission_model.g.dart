// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'mission_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$MissionModelImpl _$$MissionModelImplFromJson(Map<String, dynamic> json) =>
    _$MissionModelImpl(
      id: json['id'] as String,
      description: json['description'] as String,
      term: json['term'] as String,
      target: json['target'] as String,
      amount: json['amount'] as int,
      unit: json['unit'] as String,
      rewards: (json['rewards'] as List<dynamic>)
          .map((e) => ItemModel.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$$MissionModelImplToJson(_$MissionModelImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'description': instance.description,
      'term': instance.term,
      'target': instance.target,
      'amount': instance.amount,
      'unit': instance.unit,
      'rewards': instance.rewards,
    };
