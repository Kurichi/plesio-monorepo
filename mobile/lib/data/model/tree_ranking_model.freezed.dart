// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'tree_ranking_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

TreeRankingModel _$TreeRankingModelFromJson(Map<String, dynamic> json) {
  return _TreeRankingModel.fromJson(json);
}

/// @nodoc
mixin _$TreeRankingModel {
  String get treeId => throw _privateConstructorUsedError;
  int get rank => throw _privateConstructorUsedError;
  double get score => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $TreeRankingModelCopyWith<TreeRankingModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TreeRankingModelCopyWith<$Res> {
  factory $TreeRankingModelCopyWith(
          TreeRankingModel value, $Res Function(TreeRankingModel) then) =
      _$TreeRankingModelCopyWithImpl<$Res, TreeRankingModel>;
  @useResult
  $Res call({String treeId, int rank, double score});
}

/// @nodoc
class _$TreeRankingModelCopyWithImpl<$Res, $Val extends TreeRankingModel>
    implements $TreeRankingModelCopyWith<$Res> {
  _$TreeRankingModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? treeId = null,
    Object? rank = null,
    Object? score = null,
  }) {
    return _then(_value.copyWith(
      treeId: null == treeId
          ? _value.treeId
          : treeId // ignore: cast_nullable_to_non_nullable
              as String,
      rank: null == rank
          ? _value.rank
          : rank // ignore: cast_nullable_to_non_nullable
              as int,
      score: null == score
          ? _value.score
          : score // ignore: cast_nullable_to_non_nullable
              as double,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$TreeRankingModelImplCopyWith<$Res>
    implements $TreeRankingModelCopyWith<$Res> {
  factory _$$TreeRankingModelImplCopyWith(_$TreeRankingModelImpl value,
          $Res Function(_$TreeRankingModelImpl) then) =
      __$$TreeRankingModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String treeId, int rank, double score});
}

/// @nodoc
class __$$TreeRankingModelImplCopyWithImpl<$Res>
    extends _$TreeRankingModelCopyWithImpl<$Res, _$TreeRankingModelImpl>
    implements _$$TreeRankingModelImplCopyWith<$Res> {
  __$$TreeRankingModelImplCopyWithImpl(_$TreeRankingModelImpl _value,
      $Res Function(_$TreeRankingModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? treeId = null,
    Object? rank = null,
    Object? score = null,
  }) {
    return _then(_$TreeRankingModelImpl(
      treeId: null == treeId
          ? _value.treeId
          : treeId // ignore: cast_nullable_to_non_nullable
              as String,
      rank: null == rank
          ? _value.rank
          : rank // ignore: cast_nullable_to_non_nullable
              as int,
      score: null == score
          ? _value.score
          : score // ignore: cast_nullable_to_non_nullable
              as double,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$TreeRankingModelImpl extends _TreeRankingModel {
  const _$TreeRankingModelImpl(
      {required this.treeId, required this.rank, required this.score})
      : super._();

  factory _$TreeRankingModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$TreeRankingModelImplFromJson(json);

  @override
  final String treeId;
  @override
  final int rank;
  @override
  final double score;

  @override
  String toString() {
    return 'TreeRankingModel(treeId: $treeId, rank: $rank, score: $score)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$TreeRankingModelImpl &&
            (identical(other.treeId, treeId) || other.treeId == treeId) &&
            (identical(other.rank, rank) || other.rank == rank) &&
            (identical(other.score, score) || other.score == score));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, treeId, rank, score);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$TreeRankingModelImplCopyWith<_$TreeRankingModelImpl> get copyWith =>
      __$$TreeRankingModelImplCopyWithImpl<_$TreeRankingModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$TreeRankingModelImplToJson(
      this,
    );
  }
}

abstract class _TreeRankingModel extends TreeRankingModel {
  const factory _TreeRankingModel(
      {required final String treeId,
      required final int rank,
      required final double score}) = _$TreeRankingModelImpl;
  const _TreeRankingModel._() : super._();

  factory _TreeRankingModel.fromJson(Map<String, dynamic> json) =
      _$TreeRankingModelImpl.fromJson;

  @override
  String get treeId;
  @override
  int get rank;
  @override
  double get score;
  @override
  @JsonKey(ignore: true)
  _$$TreeRankingModelImplCopyWith<_$TreeRankingModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
