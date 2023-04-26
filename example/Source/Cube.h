// Fill out your copyright notice in the Description page of Project Settings.

#pragma once

#include "CoreMinimal.h"
#include "Cube.generated.h"
#include "GameFramework/Actor.h"

UCLASS()
class UNREALLINTER_API ACube : public AActor {
  GENERATED_BODY()

public:
  // Sets default values for this actor's properties
  ACube();

protected:
  // Called when the game starts or when spawned
  virtual void BeginPlay() override;

public:
  // Called every frame
  virtual void Tick(float DeltaTime) override;

  void Cleanup();
  void Heavy();
};
