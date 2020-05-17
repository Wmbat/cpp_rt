cmake_minimum_required(VERSION 3.14 FATAL_ERROR)
project(cpp_pt VERSION 0.0.1 DESCRIPTION "A Vulkan based rendering engine" LANGUAGES CXX)

if (NOT CMAKE_BUILD_TYPE)
   set(CMAKE_BUILD_TYPE Release)
endif()

message(STATUS "[${PROJECT_NAME}] Compiling with ${CMAKE_CXX_COMPILER_ID}")

set(CMAKE_EXPORT_COMPILE_COMMANDS ON CACHE BOOL "" FORCE)

find_package(SDL2 REQUIRED)

add_executable(${PROJECT_NAME})

set_target_properties(${PROJECT_NAME} PROPERTIES CXX_EXTENSIONS OFF)

target_compile_features(${PROJECT_NAME} PRIVATE cxx_std_20)

target_link_libraries(${PROJECT_NAME} PRIVATE SDL2 SDL2_image -pthread)

target_compile_options(${PROJECT_NAME}
   PRIVATE
      $<$<PLATFORM_ID:UNIX>:-pthread>

      $<$<AND:$<CXX_COMPILER_ID:Clang>,$<CONFIG:DEBUG>>:-o0 -Wall -Wextra>
      $<$<AND:$<CXX_COMPILER_ID:Clang>,$<CONFIG:RELEASE>>:-o3>

      $<$<AND:$<CXX_COMPILER_ID:GNU>,$<CONFIG:DEBUG>>:-o0 -g -Wall -Wextra>
      $<$<AND:$<CXX_COMPILER_ID:GNU>,$<CONFIG:RELEASE>>:-o3>
)

target_include_directories(${PROJECT_NAME}
   PUBLIC 
      $<INSTALL_INTERFACE:include>    
      $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
   PRIVATE
      ${CMAKE_CURRENT_SOURCE_DIR}/source
      ${SDL2_INCLUDE_DIRS}
)

target_sources(${PROJECT_NAME}
   PRIVATE 
      source/details/image.cpp
      source/details/pixel.cpp
      source/details/progress_keeper.cpp

      source/materials/dielectric.cpp
      source/materials/diffuse.cpp
      source/materials/metallic.cpp

      source/math/ortho_normal_basis.cpp
      source/math/vec.cpp

      source/camera.cpp
      source/main.cpp
      source/ray.cpp
      source/scene.cpp
)
