#!/usr/local/bin/perl

use strict;
use warnings;
use AnyEvent;
use AnyEvent::HTTP;

my $cocurrent = 1000;    # 并发数
my @todoList = map { "keyword" . $_ } (1 .. 50000); # 待查询的关键词

my $cv = AnyEvent->condvar;

doit() foreach 1..$cocurrent;

sub doit{
    my $word = shift @todoList;
    return if not defined $word;

    $cv->begin;
    #http_get( "http://www.baidu.com", sub { done( "http://www.baidu.com", @_ ) } );
    #http_get("http://192.168.6.151:2003", sub { done( "http://192.168.6.151:2003", @_ ) } );
    http_get("http://192.168.6.151:80/", sub { done( "http://192.168.6.151:80/", @_ ) } );
}

sub done {
    my ($word, $content, $hdr) = @_;

    $cv->end();
    print "Search: $word\tStatus: ", $hdr->{Status}, "\n";
    doit();
}

$cv->recv();