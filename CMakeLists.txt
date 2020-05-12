cmake_minimum_required(VERSION 3.14 FATAL_ERROR)
project(cpp_rt VERSION 0.0.1 DESCRIPTION "A Vulkan based rendering engine" LANGUAGES CXX)

if (NOT CMAKE_BUILD_TYPE)
   set(CMAKE_BUILD_TYPE Release)
endif()

message(STATUS "[${PROJECT_NAME}] Compiling with ${CMAKE_CXX_COMPILER_ID}")

set(CMAKE_EXPORT_COMPILE_COMMANDS ON CACHE BOOL "" FORCE)

include(FetchContent)

FetchContent_Declare(glm
   GIT_REPOSITORY https://github.com/g-truc/glm
   GIT_TAG 0.9.9.8
)

message(STATUS "[${PROJECT_NAME}] Fetching glm 0.9.9.8")

FetchContent_MakeAvailable(glm)

add_executable(${PROJECT_NAME})

set_target_properties(${PROJECT_NAME} PROPERTIES CXX_EXTENSIONS OFF)

target_compile_features(${PROJECT_NAME} PRIVATE cxx_std_20)

target_link_libraries(${PROJECT_NAME}
   PRIVATE
      glm
)

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
)

target_sources(${PROJECT_NAME}
   PRIVATE 
      source/main.cpp
      source/ray.cpp
)
