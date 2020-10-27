set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

cmake_minimum_required (VERSION ${CMAKE_VERSION})
include_directories(${CMAKE_CURRENT_BINARY_DIR}/include)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/include)

project(http_server C CXX Go)
include(cmake/golang.cmake)
include(cmake/flags.cmake)
set(SOURCE_FILE main.go)
add_subdirectory(cal)

#set(CGO_LD_FLAGS cgo -L /lib64 -lcal-api)

#set(CMAKE_EXE_LINKER_FLAGS ${CGO_LD_FLAGS})



add_go_executable(http_server ${SOURCE_FILE})

add_dependencies(http_server cal-api)

#install(TARGETS http_server DESTINATION bin COMPONENT cal-test)
