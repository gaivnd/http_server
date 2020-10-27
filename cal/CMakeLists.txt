# Generate shared library by default.
set(CAL_NAME "cal-api")
# header files.

include_directories(${CMAKE_CURRENT_BINARY_DIR}/../include)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../include)
# Collect all source files.

#set(CMAKE_C_COMPILER /usr/bin/gcc)
set(SRC
    cal.cpp
   )

add_library (${CAL_NAME} SHARED ${SRC} )
#target_include_directories(${CAL_NAME} PRIVATE ../include/)
#target_link_directories(${CAL_NAME} PRIVATE /opt/LCP/obj/centos_x86-64/sde/opt/LU3P/lib64)
target_compile_options(${CAL_NAME} PRIVATE ${CMAKE_C++_FLAGS})
#set_target_properties(${CAL_NAME} PROPERTIES LIBRARY_OUTPUT_NAME ${CAL_NAME})
#set_target_properties(${CAL_NAME} PROPERTIES LIBRARY_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/services)
#TARGET_LINK_LIBRARIES(${CAL_NAME} libgrpc++_unsecure.a protobuf libgrpc.a)

set(CAL_LOC /lib64)
#install(TARGETS ${CAL_NAME}
#DESTINATION ${CAL_LOC}
#    COMPONENT cal-api
#)
install(FILES ${CMAKE_BINARY_DIR}/libcal-api.so DESTINATION /usr/local/lib64 COMPONENT cal-api)
