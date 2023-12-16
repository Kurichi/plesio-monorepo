// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'mission_model.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

MissionModel _$MissionModelFromJson(Map<String, dynamic> json) {
  return _MissionModel.fromJson(json);
}

/// @nodoc
mixin _$MissionModel {
  String get id => throw _privateConstructorUsedError;
  String get description => throw _privateConstructorUsedError;
  String get term => throw _privateConstructorUsedError;
  String get target => throw _privateConstructorUsedError;
  int get amount => throw _privateConstructorUsedError;
  String get unit => throw _privateConstructorUsedError;
  List<ItemModel> get rewards => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $MissionModelCopyWith<MissionModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $MissionModelCopyWith<$Res> {
  factory $MissionModelCopyWith(
          MissionModel value, $Res Function(MissionModel) then) =
      _$MissionModelCopyWithImpl<$Res, MissionModel>;
  @useResult
  $Res call(
      {String id,
      String description,
      String term,
      String target,
      int amount,
      String unit,
      List<ItemModel> rewards});
}

/// @nodoc
class _$MissionModelCopyWithImpl<$Res, $Val extends MissionModel>
    implements $MissionModelCopyWith<$Res> {
  _$MissionModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? description = null,
    Object? term = null,
    Object? target = null,
    Object? amount = null,
    Object? unit = null,
    Object? rewards = null,
  }) {
    return _then(_value.copyWith(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as String,
      description: null == description
          ? _value.description
          : description // ignore: cast_nullable_to_non_nullable
              as String,
      term: null == term
          ? _value.term
          : term // ignore: cast_nullable_to_non_nullable
              as String,
      target: null == target
          ? _value.target
          : target // ignore: cast_nullable_to_non_nullable
              as String,
      amount: null == amount
          ? _value.amount
          : amount // ignore: cast_nullable_to_non_nullable
              as int,
      unit: null == unit
          ? _value.unit
          : unit // ignore: cast_nullable_to_non_nullable
              as String,
      rewards: null == rewards
          ? _value.rewards
          : rewards // ignore: cast_nullable_to_non_nullable
              as List<ItemModel>,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$MissionModelImplCopyWith<$Res>
    implements $MissionModelCopyWith<$Res> {
  factory _$$MissionModelImplCopyWith(
          _$MissionModelImpl value, $Res Function(_$MissionModelImpl) then) =
      __$$MissionModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String id,
      String description,
      String term,
      String target,
      int amount,
      String unit,
      List<ItemModel> rewards});
}

/// @nodoc
class __$$MissionModelImplCopyWithImpl<$Res>
    extends _$MissionModelCopyWithImpl<$Res, _$MissionModelImpl>
    implements _$$MissionModelImplCopyWith<$Res> {
  __$$MissionModelImplCopyWithImpl(
      _$MissionModelImpl _value, $Res Function(_$MissionModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? id = null,
    Object? description = null,
    Object? term = null,
    Object? target = null,
    Object? amount = null,
    Object? unit = null,
    Object? rewards = null,
  }) {
    return _then(_$MissionModelImpl(
      id: null == id
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as String,
      description: null == description
          ? _value.description
          : description // ignore: cast_nullable_to_non_nullable
              as String,
      term: null == term
          ? _value.term
          : term // ignore: cast_nullable_to_non_nullable
              as String,
      target: null == target
          ? _value.target
          : target // ignore: cast_nullable_to_non_nullable
              as String,
      amount: null == amount
          ? _value.amount
          : amount // ignore: cast_nullable_to_non_nullable
              as int,
      unit: null == unit
          ? _value.unit
          : unit // ignore: cast_nullable_to_non_nullable
              as String,
      rewards: null == rewards
          ? _value._rewards
          : rewards // ignore: cast_nullable_to_non_nullable
              as List<ItemModel>,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$MissionModelImpl extends _MissionModel {
  const _$MissionModelImpl(
      {required this.id,
      required this.description,
      required this.term,
      required this.target,
      required this.amount,
      required this.unit,
      required final List<ItemModel> rewards})
      : _rewards = rewards,
        super._();

  factory _$MissionModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$MissionModelImplFromJson(json);

  @override
  final String id;
  @override
  final String description;
  @override
  final String term;
  @override
  final String target;
  @override
  final int amount;
  @override
  final String unit;
  final List<ItemModel> _rewards;
  @override
  List<ItemModel> get rewards {
    if (_rewards is EqualUnmodifiableListView) return _rewards;
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_rewards);
  }

  @override
  String toString() {
    return 'MissionModel(id: $id, description: $description, term: $term, target: $target, amount: $amount, unit: $unit, rewards: $rewards)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$MissionModelImpl &&
            (identical(other.id, id) || other.id == id) &&
            (identical(other.description, description) ||
                other.description == description) &&
            (identical(other.term, term) || other.term == term) &&
            (identical(other.target, target) || other.target == target) &&
            (identical(other.amount, amount) || other.amount == amount) &&
            (identical(other.unit, unit) || other.unit == unit) &&
            const DeepCollectionEquality().equals(other._rewards, _rewards));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, id, description, term, target,
      amount, unit, const DeepCollectionEquality().hash(_rewards));

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$MissionModelImplCopyWith<_$MissionModelImpl> get copyWith =>
      __$$MissionModelImplCopyWithImpl<_$MissionModelImpl>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$MissionModelImplToJson(
      this,
    );
  }
}

abstract class _MissionModel extends MissionModel {
  const factory _MissionModel(
      {required final String id,
      required final String description,
      required final String term,
      required final String target,
      required final int amount,
      required final String unit,
      required final List<ItemModel> rewards}) = _$MissionModelImpl;
  const _MissionModel._() : super._();

  factory _MissionModel.fromJson(Map<String, dynamic> json) =
      _$MissionModelImpl.fromJson;

  @override
  String get id;
  @override
  String get description;
  @override
  String get term;
  @override
  String get target;
  @override
  int get amount;
  @override
  String get unit;
  @override
  List<ItemModel> get rewards;
  @override
  @JsonKey(ignore: true)
  _$$MissionModelImplCopyWith<_$MissionModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

UserMissionModel _$UserMissionModelFromJson(Map<String, dynamic> json) {
  return _UserMissionModel.fromJson(json);
}

/// @nodoc
mixin _$UserMissionModel {
  String get userId => throw _privateConstructorUsedError;
  MissionModel get mission => throw _privateConstructorUsedError;
  int get progress => throw _privateConstructorUsedError;
  int get deadline => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $UserMissionModelCopyWith<UserMissionModel> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $UserMissionModelCopyWith<$Res> {
  factory $UserMissionModelCopyWith(
          UserMissionModel value, $Res Function(UserMissionModel) then) =
      _$UserMissionModelCopyWithImpl<$Res, UserMissionModel>;
  @useResult
  $Res call({String userId, MissionModel mission, int progress, int deadline});

  $MissionModelCopyWith<$Res> get mission;
}

/// @nodoc
class _$UserMissionModelCopyWithImpl<$Res, $Val extends UserMissionModel>
    implements $UserMissionModelCopyWith<$Res> {
  _$UserMissionModelCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? userId = null,
    Object? mission = null,
    Object? progress = null,
    Object? deadline = null,
  }) {
    return _then(_value.copyWith(
      userId: null == userId
          ? _value.userId
          : userId // ignore: cast_nullable_to_non_nullable
              as String,
      mission: null == mission
          ? _value.mission
          : mission // ignore: cast_nullable_to_non_nullable
              as MissionModel,
      progress: null == progress
          ? _value.progress
          : progress // ignore: cast_nullable_to_non_nullable
              as int,
      deadline: null == deadline
          ? _value.deadline
          : deadline // ignore: cast_nullable_to_non_nullable
              as int,
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $MissionModelCopyWith<$Res> get mission {
    return $MissionModelCopyWith<$Res>(_value.mission, (value) {
      return _then(_value.copyWith(mission: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$UserMissionModelImplCopyWith<$Res>
    implements $UserMissionModelCopyWith<$Res> {
  factory _$$UserMissionModelImplCopyWith(_$UserMissionModelImpl value,
          $Res Function(_$UserMissionModelImpl) then) =
      __$$UserMissionModelImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String userId, MissionModel mission, int progress, int deadline});

  @override
  $MissionModelCopyWith<$Res> get mission;
}

/// @nodoc
class __$$UserMissionModelImplCopyWithImpl<$Res>
    extends _$UserMissionModelCopyWithImpl<$Res, _$UserMissionModelImpl>
    implements _$$UserMissionModelImplCopyWith<$Res> {
  __$$UserMissionModelImplCopyWithImpl(_$UserMissionModelImpl _value,
      $Res Function(_$UserMissionModelImpl) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? userId = null,
    Object? mission = null,
    Object? progress = null,
    Object? deadline = null,
  }) {
    return _then(_$UserMissionModelImpl(
      userId: null == userId
          ? _value.userId
          : userId // ignore: cast_nullable_to_non_nullable
              as String,
      mission: null == mission
          ? _value.mission
          : mission // ignore: cast_nullable_to_non_nullable
              as MissionModel,
      progress: null == progress
          ? _value.progress
          : progress // ignore: cast_nullable_to_non_nullable
              as int,
      deadline: null == deadline
          ? _value.deadline
          : deadline // ignore: cast_nullable_to_non_nullable
              as int,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$UserMissionModelImpl extends _UserMissionModel {
  const _$UserMissionModelImpl(
      {required this.userId,
      required this.mission,
      required this.progress,
      required this.deadline})
      : super._();

  factory _$UserMissionModelImpl.fromJson(Map<String, dynamic> json) =>
      _$$UserMissionModelImplFromJson(json);

  @override
  final String userId;
  @override
  final MissionModel mission;
  @override
  final int progress;
  @override
  final int deadline;

  @override
  String toString() {
    return 'UserMissionModel(userId: $userId, mission: $mission, progress: $progress, deadline: $deadline)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$UserMissionModelImpl &&
            (identical(other.userId, userId) || other.userId == userId) &&
            (identical(other.mission, mission) || other.mission == mission) &&
            (identical(other.progress, progress) ||
                other.progress == progress) &&
            (identical(other.deadline, deadline) ||
                other.deadline == deadline));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode =>
      Object.hash(runtimeType, userId, mission, progress, deadline);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$UserMissionModelImplCopyWith<_$UserMissionModelImpl> get copyWith =>
      __$$UserMissionModelImplCopyWithImpl<_$UserMissionModelImpl>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$UserMissionModelImplToJson(
      this,
    );
  }
}

abstract class _UserMissionModel extends UserMissionModel {
  const factory _UserMissionModel(
      {required final String userId,
      required final MissionModel mission,
      required final int progress,
      required final int deadline}) = _$UserMissionModelImpl;
  const _UserMissionModel._() : super._();

  factory _UserMissionModel.fromJson(Map<String, dynamic> json) =
      _$UserMissionModelImpl.fromJson;

  @override
  String get userId;
  @override
  MissionModel get mission;
  @override
  int get progress;
  @override
  int get deadline;
  @override
  @JsonKey(ignore: true)
  _$$UserMissionModelImplCopyWith<_$UserMissionModelImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
