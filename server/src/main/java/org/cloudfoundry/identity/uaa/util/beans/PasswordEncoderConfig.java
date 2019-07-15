package org.cloudfoundry.identity.uaa.util.beans;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.DelegatingPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;

import java.util.HashMap;
import java.util.Map;

@Configuration
public class PasswordEncoderConfig {

    private static Logger logger = LoggerFactory.getLogger(PasswordEncoderConfig.class);

    @Bean
    public PasswordEncoder nonCachingPasswordEncoder() {
        logger.info("Building DelegatingPasswordEncoder with {bcrypt}");

        Map<String, PasswordEncoder> encoders = new HashMap<>();

        logger.info("Building DelegatingPasswordEncoder : Adding \"bcrypt\" to encoders");
        encoders.put("bcrypt", new BCryptPasswordEncoder());

        logger.info("Building DelegatingPasswordEncoder : Adding \"null\" to encoders");
        encoders.put("null", new BCryptPasswordEncoder());

        logger.info("Building DelegatingPasswordEncoder : Adding null to encoders");
        encoders.put(null, new BCryptPasswordEncoder());
        return new DelegatingPasswordEncoder("bcrypt", encoders);
    }
}
